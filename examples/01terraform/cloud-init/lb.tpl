#cloud-config

# set locale
locale: en_US.UTF-8

# set timezone
# timezone: Etc/UTC

# set root password
chpasswd:
  list: |
    root:linux
    ${username}:${password}
  expire: False

ssh_authorized_keys:
${authorized_keys}

ntp:
  enabled: true
  ntp_client: chrony
  config:
    confpath: /etc/chrony.conf
  servers:
${ntp_servers}

# need to disable gpg checks because the cloud image has an untrusted repo
zypper:
  repos:
${repositories}
  config:
    gpgcheck: "off"
    solver.onlyRequires: "true"
    download.use_deltarpm: "true"

# need to remove the standard docker packages that are pre-installed on the
# cloud image because they conflict with the kubic- ones that are pulled by
# the kubernetes packages
packages:
  - haproxy

write_files:
- path: /etc/haproxy/haproxy.cfg
  content: |
    defaults
      timeout connect 10s
      timeout client 86400s
      timeout server 86400s

    listen stats
      bind    *:9000
      mode    http
      stats   hide-version
      stats   uri       /stats

    frontend apiserver
      bind :6443
      default_backend apiserver-backend

    frontend gangway
      bind :32001
      default_backend gangway-backend

    frontend dex
      bind :32000
      default_backend dex-backend

    backend apiserver-backend
      option httpchk GET /healthz
      ${apiserver_backends}

    backend gangway-backend
      option httpchk GET /
      ${gangway_backends}

    backend dex-backend
      option httpchk GET /healthz
      ${dex_backends}

runcmd:
  # Since we are currently inside of the cloud-init systemd unit, trying to
  # start another service by either `enable --now` or `start` will create a
  # deadlock. Instead, we have to use the `--no-block-` flag.
  - /usr/bin/localectl set-keymap de-latin1-nodeadkeys
  - [ systemctl, enable, --now, --no-block, haproxy ]
  - [ systemctl, disable, --now, --no-block, firewalld ]
  # The template machine should have been cleaned up, so no machine-id exists
  - [ dbus-uuidgen, --ensure ]
  - [ systemd-machine-id-setup ]
  # With a new machine-id generated the journald daemon will work and can be restarted
  # Without a new machine-id it should be in a failed state
  - [ systemctl, restart, systemd-journald ]

bootcmd:
  - ip link set dev eth0 mtu 1400
  # Hostnames from DHCP - otherwise localhost will be used
  - /usr/bin/sed -ie "s#DHCLIENT_SET_HOSTNAME=\"no\"#DHCLIENT_SET_HOSTNAME=\"yes\"#" /etc/sysconfig/network/dhcp
  - netconfig update -f

final_message: "The system is finally up, after $UPTIME seconds"

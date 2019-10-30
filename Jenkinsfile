pipeline {
  agent any
  stages {
    stage('infra') {
      steps {
        echo 'deploy infra'
      }
    }
    stage('kubernetes') {
      steps {
        echo 'deploy kubernetes'
      }
    }
    stage('rook') {
      steps {
        echo 'deploy rook'
      }
    }
    stage('run tests') {
      steps {
        echo 'run tests'
      }
    }
  }
}
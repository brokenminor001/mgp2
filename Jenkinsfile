pipeline {
  
  agent any
  tools { go '1.20' }
  stages {
    stage('Build') {
      steps {
        sh 'git clone https://github.com/brokenminor001/mgp2.git'
        sh 'go version'
        sh 'docker ps'
        sh 'ls'
      }
      
    }
  }
}
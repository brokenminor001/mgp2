pipeline {
  
  agent any
  tools { go '1.20' }
  stages {
    stage('Build') {
      steps {
        
        sh 'go version'
        sh 'docker ps'
      }
      
    }
  }
}
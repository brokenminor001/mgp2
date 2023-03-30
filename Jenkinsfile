pipeline {
  environment {
    IMAGE_BASE = 'asmolin/mgp'
    IMAGE_TAG = "1.0.0"
    IMAGE_NAME = "${env.IMAGE_BASE}:${env.IMAGE_TAG}"
    IMAGE_NAME_LATEST = "${env.IMAGE_BASE}:latest"
    DOCKERFILE_NAME = "Dockerfile"

  }
  agent any
  tools { go '1.20' }
  stages {
    stage('Build') {
      steps {
        checkout scm
        
        sh 'go build main.go'
        
      }
    }
       stage('Push images') {
      agent any
      when {
        branch 'master'
      }
      steps {
        script {
          def dockerImage = docker.build("${env.IMAGE_NAME}", "-f ${env.DOCKERFILE_NAME} .")
          docker.withRegistry('', 'dockerhub') {
            dockerImage.push()
            dockerImage.push("latest")
          }
          echo "Pushed Docker Image: ${env.IMAGE_NAME}"
        }
        sh "docker rmi ${env.IMAGE_NAME} ${env.IMAGE_NAME_LATEST}"
      }
    }
      
  }
  }

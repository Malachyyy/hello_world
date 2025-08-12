pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build') {
            agent {
                docker {
                    image 'golang:tip-alpine3.22'
                    args '-v $WORKSPACE:/go/src/hello_world'
                }
            }
            steps {
                sh 'go build -o hello_world main.go'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
    }
}

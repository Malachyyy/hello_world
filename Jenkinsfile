pipeline {
    agent {
        docker { image 'golang:1.24.6' }
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build') {
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

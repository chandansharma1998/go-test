pipeline {
    agent any

    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-creds')
        DOCKERHUB_USER = 'sharmachandan487'
        IMAGE_NAME = 'go-test'
    }

    stages {
        stage('Source') {
            steps {
                git branch: 'main', url: 'https://github.com/chandansharma1998/go-test.git'
            }
        }
        stage('Docker build and push') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', 'dockerhub-creds') {
                        def app = docker.build("${DOCKERHUB_USER}/${IMAGE_NAME}:${env.BUILD_NUMBER}")
                        app.push()
                    }
                }
            }
        }
        stage('Deploy') {
            steps {
                sh '''
                helm upgrade --install go-test ./helm-chart/go-test \
                --set image.repository=${DOCKERHUB_USER}/${IMAGE_NAME} \
                --set image.tag=${BUILD_NUMBER}
                '''
            }
        }
    }

    post {
        success {
            mail to: "${env.GIT_COMMIT_AUTHOR_EMAIL}",
                subject: "Pipeline Success: #${env.BUILD_NUMBER}",
                body: "Build and deployment succeeded!"
        }
        failure {
            mail to: "${env.GIT_COMMIT_AUTHOR_EMAIL}",
                subject: "Pipeline Failed: #${env.BUILD_NUMBER}",
                body: "Build or deployment failed. Please check Jenkins logs."
        }
    }
}
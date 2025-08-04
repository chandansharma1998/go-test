pipeline {
    agent any

    environment {
        IMAGE_NAME = 'go-test'
        DOCKERHUB_USER = 'sharmachandan487'
    }

    stages {
        stage('Source') {
            steps {
                git branch: 'main', url: 'https://github.com/chandansharma1998/go-test.git'
            }
        }

        stage('Docker build and push') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKERHUB_USER', passwordVariable: 'DOCKERHUB_PASS')]) {
                    bat '''
                        docker --version
                        docker login -u %DOCKERHUB_USER% -p %DOCKERHUB_PASS%
                        docker build -t %DOCKERHUB_USER%/%IMAGE_NAME%:%BUILD_NUMBER% .
                        docker push %DOCKERHUB_USER%/%IMAGE_NAME%:%BUILD_NUMBER%
                    '''
                }
            }
        }

        stage('Deploy') {
            steps {
                bat '''
                    helm upgrade --install go-test ./helm-chart/go-test ^
                    --set image.repository=%DOCKERHUB_USER%/%IMAGE_NAME% ^
                    --set image.tag=%BUILD_NUMBER%
                '''
            }
        }

        stage('Extract Commit Email') {
            steps {
                script {
                    def email = bat(
                        script: 'git log -1 --pretty=format:"%%ae"',
                        returnStdout: true
                    ).trim()
                    env.COMMIT_EMAIL = email
                }
            }
        }
    }

    post {
        success {
            mail to: "${env.COMMIT_EMAIL}",
                subject: "Pipeline Success: #${env.BUILD_NUMBER}",
                body: "Build and deployment succeeded!"
        }
        failure {
            mail to: "${env.COMMIT_EMAIL}",
                subject: "Pipeline Failed: #${env.BUILD_NUMBER}",
                body: "Build or deployment failed. Please check Jenkins logs."
        }
    }
}
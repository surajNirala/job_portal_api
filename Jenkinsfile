pipeline {
    agent any 

    environment {
        DOCKER_HUB_REPO = 'snirala1995/job_portal_api'
        DOCKER_IMAGE_TAG = "${DOCKER_HUB_REPO}:${env.BUILD_NUMBER}"
        CONTAINER_NAME = 'job_portal_api' // Name for your Docker container
        CONTAINER_PORT = '8080' // Port inside the Docker container
        // ENV_JOB_PORTAL_API = credentials('ENV_JOB_PORTAL_API') // For secret file
        CREDENTIAL_SNIRALA_DOCKERHUB = 'credentials-snirala-dockerhub'
        CREDENTIALS_GOLANG_SERVER = 'credentials-golang-server'
        JENKINS_SERVER = '35.200.176.111'
        GOLANG_SERVER = '34.131.166.50'
        ENV_FINAL_LIVE = '/home/srj/env/.env:/app/.env'
    }

    parameters {
        choice(name: 'ENVIRONMENT', choices: ['Dev', 'Live'], description: 'Select The Environment')
    }

    stages {
        stage('Set Ports') {
            steps {
                script {
                    // Set ports based on the selected environment
                    if (params.ENVIRONMENT == 'Dev') {
                        env.HOST_PORT = '8083'
                        env.SERVER_IP = "http://${GOLANG_SERVER}:${env.HOST_PORT}"
                    } else if (params.ENVIRONMENT == 'Live') {
                        env.HOST_PORT = '8084'
                        env.SERVER_IP = "http://${GOLANG_SERVER}:${env.HOST_PORT}"
                    }
                }
            }
        }
       /*  stage('Code Analysis') {
            environment {
                scannerHome = tool 'Sonar-Scanner'
            }
            steps {
                script {
                    withSonarQubeEnv('Sonar-Scanner') {
                        sh """
                            ${scannerHome}/bin/sonar-scanner \
                            -Dsonar.projectKey=workwebui \
                            -Dsonar.sources=. \
                            -Dsonar.host.url=https://sonarqube.dhimalu.xyz \
                            -Dsonar.login=sqp_c6b396ecc795e7b4e16eb2a48b015d326baf1477
                        """
                    }
                }
            }
        } */
        stage('Check Existing Container') {
            steps {
                script {
                    echo "Checking if the container already exists"
                    def existingContainer = sh(script: "docker ps -aqf name=${CONTAINER_NAME}-${env.HOST_PORT}", returnStdout: true).trim()
                    if (existingContainer) {
                        echo "Stopping and removing the existing container: ${CONTAINER_NAME}-${env.HOST_PORT}"
                        sh "docker rm -f ${CONTAINER_NAME}-${env.HOST_PORT}"
                    }
                }
            }
        }

        // stage('Prepare .env File') {
        //     steps {
        //         echo "Removing the existing .env file if it exists"
        //         sh 'rm -f .env'
        //         echo "Copying the new .env file"
        //         sh "cp ${ENV_JOB_PORTAL_API} .env"
        //     }
        // }

        stage('Build Docker Image') {
            steps {
                script {
                    echo "Building the Docker image"
                    docker.build(DOCKER_IMAGE_TAG)
                    echo "Docker image built successfully."
                }
            }
        }

        stage('Push Docker Image To Docker Hub') {
            steps {
                script {
                    try {
                        echo "Pushing Docker image to DockerHub."
                        docker.withRegistry('https://registry.hub.docker.com', CREDENTIAL_SNIRALA_DOCKERHUB) {
                            docker.image(DOCKER_IMAGE_TAG).push()
                        }
                        echo "Docker image pushed to DockerHub successfully."
                    } catch (Exception e) {
                        echo "Failed to push Docker image: ${e.message}"
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    if (params.ENVIRONMENT == 'Dev' || params.ENVIRONMENT == 'Live') {
                        echo "Deploying to ================= SRJ-SERVER ============== (${GOLANG_SERVER})"
                        sshagent([CREDENTIALS_GOLANG_SERVER]) {
                            echo "Deploying to ${GOLANG_SERVER} on port ${HOST_PORT} with image ${DOCKER_IMAGE_TAG}"
                            withCredentials([usernamePassword(credentialsId: CREDENTIAL_SNIRALA_DOCKERHUB, usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]){
                            sh """
                                echo "Connecting to ${GOLANG_SERVER}..."
                                ssh -o StrictHostKeyChecking=no srj@${GOLANG_SERVER} <<EOF
                                echo "Remote server connected successfully!"

                                echo "Logging into DockerHub"
                                echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

                                echo "Pulling Docker image from DockerHub: ${DOCKER_IMAGE_TAG}"
                                docker pull ${DOCKER_IMAGE_TAG}

                                echo "Stopping and removing any existing container"
                                docker rm -f ${CONTAINER_NAME}-${HOST_PORT} || true

                                echo "Running the Docker container"

                                docker run -d --init -p ${HOST_PORT}:${CONTAINER_PORT} -v ${ENV_FINAL_LIVE} --name ${CONTAINER_NAME}-${HOST_PORT} ${DOCKER_IMAGE_TAG}
                                
                                echo "Docker image ${DOCKER_IMAGE_TAG} run successfully."
                                exit
                            """
                            // docker run --env-file ${ENV_FINAL_LIVE} -d --init -p ${HOST_PORT}:${CONTAINER_PORT} --name ${CONTAINER_NAME}-${HOST_PORT} ${DOCKER_IMAGE_TAG}
                            }
                        }
                    } else {
                        echo "Deploying image in non-Dev environment"
                        sh "docker run -d --init -p ${HOST_PORT}:${CONTAINER_PORT} --name ${CONTAINER_NAME}-${HOST_PORT} ${DOCKER_IMAGE_TAG}"
                        echo "Docker image ${DOCKER_IMAGE_TAG} run successfully."
                    }   
                }
            }
        }
    }

    post {
        success {
            script {
                echo "Docker image ${DOCKER_IMAGE_TAG} successfully pushed to Docker Hub."
                echo "Container running on port: ${HOST_PORT}"
                echo "Pipeline completed successfully."
                echo "Click the following link to check the website live: ${env.SERVER_IP}"
            }
        }
        failure {
            script {
                echo "Pipeline failed. Check logs for details."
            }
        }
    }
}

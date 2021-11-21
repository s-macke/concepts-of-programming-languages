FROM openjdk:17
COPY . /usr/src/myapp
#COPY .gradle /usr/src/myapp/.gradle
WORKDIR /usr/src/myapp
RUN ./gradlew clean build -x test
CMD ["./gradlew", "bootRun"]
FROM eclipse-temurin:21-jdk-alpine AS build
COPY Main.java .
RUN javac Main.java
FROM eclipse-temurin:21-jre-alpine
COPY --from=build /Main*.class /app/
WORKDIR /app
CMD ["java","-Xmx256m","Main"]

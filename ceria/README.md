<p align="center">
    <img src="./kotlin-1.svg" width="250" height="250">
</p>

# Advent of Code 2021 - in [Kotlin](https://kotlinlang.org/)

I'm going to attempt (likely with varying degrees of motivation) to implement all of the advent of code problems in Kotlin. Since [VSCode](https://code.visualstudio.com/) and other IDEs rely on [Gradle](https://gradle.org/) or Maven at the moment to supply code completion for Kotlin, each day of the advent is a simple Gradle project. You do not need to have Gradle or Kotlin installed. Using either option of execution descbied below, you only need to have Java 11 installed.

## Running the code

I've included the input file supplied by advent of code for each problem. Both options assume you are executiing from the top level, i.e. `<day>/`.


Using gradle,
```
    ./gradlew run --args=input
```
You can substitute `input` with the path to your own input if you desire.


Alternatively, you can manually run the code,
```
   java -jar Solution.jar input
```
You may also download just this jar file and execute it from any directory `java -jar Solution.jar <path-to-input>`


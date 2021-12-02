plugins {
	application
	id("com.github.johnrengelman.shadow") version "7.1.0"
	kotlin("jvm") version "1.6.0"
}

repositories {
	mavenCentral()
}

application {
	mainClass.set("SolutionKt")
}

tasks {
	clean {
		file("${projectDir}/Solution.jar").delete()
	}

    shadowJar {
    	archiveBaseName.set("Solution")
		archiveClassifier.set("")
		destinationDirectory.set(projectDir)
    	mergeServiceFiles()
    }
}

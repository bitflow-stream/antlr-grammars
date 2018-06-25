ELK_SERVER?=cringly@elk
export CLASSPATH=".:/usr/local/lib/antlr-4.7.1-complete.jar:$CLASSPATH"
antlr4 = java -jar /usr/local/lib/antlr-4.7.1-complete.jar
grun = java org.antlr.v4.gui.TestRig
out = ./out
src = src/main/antlr/ch/chrisport/bitflow

compile-java:
	rm -rf out
	mkdir -p out
	cp src/main/antlr/ch/chrisport/bitflow/* out
	cd out && ${antlr4} Bitflow.g4 && javac Bitflow*.java

test: compile-java
	cd out && ${grun} Bitflow pipeline -token
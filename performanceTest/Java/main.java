package com.mycompany.multiplicacionmatrices;

/**
 *
 * @author Alón
 */
public class main {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        String filename = "/Users/aliwsky/NetBeansProjects/MultiplicacionMatrices/src/main/java/com/mycompany/multiplicacionmatrices/matricesToCalculate.csv";
        int nrOfThreads = 1000;
        Java_Thread java_Thread = new Java_Thread();
        java_Thread.start(filename, nrOfThreads);
    }

}

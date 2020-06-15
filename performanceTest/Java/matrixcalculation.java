package com.mycompany.multiplicacionmatrices;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.LinkedList;
import java.util.List;
import java.util.ArrayList;
import java.util.concurrent.ThreadLocalRandom;

public class Java_Thread {

    public List<ArrayList<ArrayList<Integer>>> read(String filename) {
        ArrayList<ArrayList<Integer>> A = new ArrayList<ArrayList<Integer>>();
        ArrayList<ArrayList<Integer>> B = new ArrayList<ArrayList<Integer>>();
        String thisLine;
        try {
            BufferedReader br = new BufferedReader(new FileReader(filename));
            while ((thisLine = br.readLine()) != null) {
                if (thisLine.trim().equals("")) {
                    break;
                } else {
                    ArrayList<Integer> line = new ArrayList<Integer>();
                    String[] lineArray = thisLine.split(",");
                    for (String number : lineArray) {
                        line.add(Integer.parseInt(number));
                    }
                    A.add(line);
                }
            }
            while ((thisLine = br.readLine()) != null) {
                ArrayList<Integer> line = new ArrayList<Integer>();
                String[] lineArray = thisLine.split(",");
                for (String number : lineArray) {
                    line.add(Integer.parseInt(number));
                }
                B.add(line);
            }
            br.close();
        } catch (IOException e) {
            System.err.println("Error: " + e);
        }
        List<ArrayList<ArrayList<Integer>>> res = new LinkedList<ArrayList<ArrayList<Integer>>>();
        res.add(A);
        res.add(B);
        return res;
    }

    public List<ArrayList<ArrayList<Integer>>> create() {
        ArrayList<ArrayList<Integer>> A = new ArrayList<ArrayList<Integer>>();
        ArrayList<ArrayList<Integer>> B = new ArrayList<ArrayList<Integer>>();
        String thisLine;

        int largo = 100;
        int ancho = 100;

        for (int i = 0; i < largo; i++) {
            A.add(new ArrayList<Integer>());
            B.add(new ArrayList<Integer>());
            for (int j = 0; j < ancho; j++) {
                A.get(A.size()-1).add(ThreadLocalRandom.current().nextInt(0, 101 + 1));
                B.get(B.size()-1).add(ThreadLocalRandom.current().nextInt(0, 101 + 1));
            }
        }

        List<ArrayList<ArrayList<Integer>>> res = new LinkedList<ArrayList<ArrayList<Integer>>>();
        res.add(A);
        res.add(B);
        return res;
    }

    public int[][] matrixMultiplication(ArrayList<ArrayList<Integer>> A, ArrayList<ArrayList<Integer>> B, int m, int n) {
        int[][] C = new int[m][n];
        for (int i = 0; i < m; i++) {
            for (int k = 0; k < n; k++) {
                int temp = A.get(i).get(k);
                for (int j = 0; j < n; j++) {
                    C[i][j] += temp * B.get(k).get(j);
                }
            }
        }
        return C;
    }

    public ArrayList<ArrayList<ArrayList<Integer>>>
            splitMatrix(ArrayList<ArrayList<Integer>> A, int nrOfThreads) {
        int n = A.size();
        int m = n / nrOfThreads;
        ArrayList<ArrayList<ArrayList<Integer>>> B = new ArrayList<ArrayList<ArrayList<Integer>>>();
        for (int i = 0; i < nrOfThreads; i++) {
            B.add(new ArrayList<ArrayList<Integer>>(A.subList(i * m,
                    (i + 1) * m)));
        }
        return B;
    }

    public void start(String filename, int nrOfThreads) {
        //Load matrices form file
        List<ArrayList<ArrayList<Integer>>> matrices = read(filename);
        //List<ArrayList<ArrayList<Integer>>> matrices = create();

        ArrayList<ArrayList<Integer>> A = matrices.get(0);
        ArrayList<ArrayList<Integer>> B = matrices.get(1);

        if (nrOfThreads <= 0) {
            //Run it sequentially
            int n = A.size();
            long startTime = System.nanoTime();
            int[][] C = matrixMultiplication(A, B, n, n);
            long endTime = System.nanoTime();

            System.out.println("Execution took " + (endTime
                    - startTime) + " ns");

        } else {
            if (A.size() % nrOfThreads != 0) {
                System.out.println(
                        "Size of matrix is not divisible by the supplied number of threads");
                System.exit(1);
            }
            ArrayList<int[][]> result = new ArrayList<int[][]>();
            int[][] empty = new int[][]{{}};
            for (int i = 0; i < nrOfThreads; i++) {
                result.add(empty);
            }

            //Split matrix for each thread
            ArrayList<ArrayList<ArrayList<Integer>>> workerMatrices = splitMatrix(A, nrOfThreads);
            ArrayList<Worker> threads = new ArrayList<Worker>();
            long startTime = System.nanoTime();
            for (int i = 0; i < nrOfThreads; i++) {
                //Give to each thread their section of matrix A, matrix B, index of the result section, and the reference to the result matrix
                threads.add(new Worker(workerMatrices.get(i), B,
                        i, result));
                //Start the process
                threads.get(i).start();
            }

            //Wait all the threads finished
            for (int i = 0; i < nrOfThreads; i++) {
                try {
                    threads.get(i).join();
                } catch (Exception e) {
                    System.err.println(e);
                }
            }
            long endTime = System.nanoTime();
            System.out.println("Execution took " + (endTime
                    - startTime) + " ns");
        }
    }

    class Worker extends Thread {

        private ArrayList<ArrayList<Integer>> A;
        private ArrayList<ArrayList<Integer>> B;
        private int index;
        private ArrayList<int[][]> result;
        private int m;
        private int n;

        public Worker(ArrayList<ArrayList<Integer>> A,
                ArrayList<ArrayList<Integer>> B, int index,
                ArrayList<int[][]> result) {
            this.A = A;
            this.B = B;
            this.index = index;
            this.result = result;
            this.m = A.size();
            this.n = B.size();
        }

        @Override
        public void run() {
            //Get the result of it's section and save it in the corresponding position of the result matrix
            this.result.set(this.index,
                    matrixMultiplication(this.A, this.B, this.m, this.n));
        }
    }
}

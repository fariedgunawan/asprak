import java.util.Scanner;

public class matrix_mult {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        // IS. program menerima input ukuran matriks n x n
        // FS. variabel n telah diinisialisasi dengan ukuran matriks yang diberikan
        System.out.print("Masukkan n x n ukuran matrix : ");
        int n = input.nextInt();

        int[][] matrix1 = new int[n][n];
        int[][] matrix2 = new int[n][n];
        int[][] matrixhasil = new int[n][n];

        // IS. matrix1 belum diinisialisasi dengan nilai-nilai yang diberikan
        // FS. matrix1 telah diinisialisasi dengan nilai-nilai yang dimasukkan oleh pengguna
        System.out.println("Masukkan Nilai Matriks ke-1 : ");
        inputMatrix(matrix1, input, matrix1.length, matrix1[0].length);

        // IS. matrix2 belum diinisialisasi dengan nilai-nilai yang diberikan
        // FS. matrix2 telah diinisialisasi dengan nilai-nilai yang dimasukkan oleh pengguna
        System.out.println("Masukkan Nilai Matriks ke-2 : ");
        inputMatrix(matrix2, input, matrix2.length, matrix2[0].length);

        input.close();

        // IS. matrix1 telah diinisialisasi
        // FS. nilai-nilai dari matrix1 telah ditampilkan ke layar
        System.out.println("Matriks ke-1:");
        printMatrix(matrix1, matrix1.length, matrix1[0].length);

        // IS. matrix2 telah diinisialisasi
        // FS. nilai-nilai dari matrix2 telah ditampilkan ke layar
        System.out.println("Matriks Ke-2:");
        printMatrix(matrix2, matrix2.length, matrix2[0].length);

        // IS. matrixhasil belum diinisialisasi
        // FS. matrixhasil telah diisi dengan hasil perkalian matrix1 dan matrix2
        multiplyMatrix(matrix1, matrix2, matrixhasil);

        // IS. matrixhasil telah diinisialisasi dengan hasil perkalian
        // FS. hasil perkalian matrixhasil telah ditampilkan ke layar
        System.out.println("Hasil Perkalian Matriks:");
        printMatrix(matrixhasil, matrixhasil.length, matrixhasil[0].length);
    }

    // IS. matrix belum diisi dengan nilai apapun
    // FS. matrix telah diisi dengan nilai-nilai yang dimasukkan oleh pengguna
    public static void inputMatrix(int[][] matrix, Scanner input, int rowLength, int colLength) {
        for (int row = 0;row < rowLength;row++){
            for (int col = 0;col < colLength;col++){
                matrix[row][col] = input.nextInt();
            }
        }
    }

    // IS. matrix telah diinisialisasi
    // FS. nilai-nilai dari matrix telah ditampilkan ke layar
    public static void printMatrix(int[][] matrix, int rowLength, int colLength) {
        for (int row = 0;row < rowLength;row++){
            for (int col = 0;col < colLength;col++){
                System.out.print(matrix[row][col]+" ");
            }
            System.out.println("");
        }
    }

    // IS. matrixhasil belum diinisialisasi
    // FS. matrixhasil telah diisi dengan hasil perkalian matrix1 dan matrix2
    public static void multiplyMatrix(int[][] matrix1, int[][] matrix2, int[][] matrixhasil) {
        for (int row = 0;row < matrix1.length;row++){
            for (int col = 0;col < matrix2[0].length;col++){
                for (int i = 0;i<matrix2.length;i++){
                    matrixhasil[row][col] += matrix1[row][i] * matrix2[i][col];
                }
            }
        }
    }
}

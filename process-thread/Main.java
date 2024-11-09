class Main {
    static boolean isPrime(long num) {
        // Corner cases
        if (num <= 1) {
            return false;
        }
        if (num <= 3) {
            return true;
        }

        // This is checked so that we can skip
        // middle five numbers in below loop
        if (num % 2l == 0 || num % 3l == 0)
            return false;

        for (long i = 5l; i * i <= num; i = i + 6l)
            if (num % i == 0 || num % (i + 2l) == 0)
                return false;

        return true;

    }

    static void findPrimesBetween(long start, long end) {
        for (var i = start; i < end; i++) {
            if (isPrime(i)) {
                System.out.printf("%d is Prime\n", i);
            }
        }
    }

    static void printNums(int threadNumber, int start, int end) {
        for (var i = start; i < end; i++) {
            try {
                Thread.sleep(100);
            } catch (Exception e) {

            }
            System.out.printf("Thread%d - %d\n", threadNumber, i);
        }

    }

    public static void main(String args[]) throws Exception {
        Thread thread1 = new Thread(() -> printNums(1, 1, 10));
        Thread thread2 = new Thread(() -> printNums(2, 1, 10));
        thread1.start();
        thread2.start();

        thread1.join();
        thread2.join();

    }

    public static void main1(String args[]) throws Exception {

        long start = 1000000000000l;
        long count = 100000l * 2l;

        long start1 = start;
        long end1 = start1 + (count / 2);

        long start2 = end1;
        long end2 = start2 + (count / 2);

        // findPrimesBetween(start1, end1);
        // findPrimesBetween(start2, end2);

        Thread thread1 = new Thread(() -> findPrimesBetween(start1, end1));
        Thread thread2 = new Thread(() -> findPrimesBetween(start2, end2));

        thread1.start();
        thread2.start();

        thread1.join();
        thread2.join();

    }
}

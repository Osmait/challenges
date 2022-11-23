public  class index {
    /*
Given a sequence of integers as an array, determine whether it is possible to obtain a strictly increasing sequence by removing no more than one element from the array.

Note: sequence a0, a1, ..., an is considered to be a strictly increasing if a0 < a1 < ... < an. Sequence containing only one element is also considered to be strictly increasing.
*/

    public static void main(String[] args) {
        int[] lists ={1, 2, 1, 2};
        solution(lists);
    }

    static boolean solution(int[] sequence) {
       int screws = 0;
       int tail = 0;
        for (int i = 1; i < sequence.length; i++) {
            if (sequence[tail] < sequence[i]) {
                tail = i;
            }else {
                if (tail == 0 || sequence[i]> sequence[tail-1]){
                    tail = i;
                }
                screws ++;
            }
        }
        return screws == 1;
    }


}
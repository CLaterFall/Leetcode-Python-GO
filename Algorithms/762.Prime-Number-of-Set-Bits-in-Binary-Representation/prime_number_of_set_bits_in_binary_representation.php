class Solution {

    /**
     * @param Integer $L
     * @param Integer $R
     * @return Integer
     */
    function countPrimeSetBits($L, $R) {
        $res = 0;
        for($i=$L; $i<=$R; $i++) {
            $binstr = decbin($i);
            $count = 0;
            $len = strlen($binstr);
            for ($j=0; $j<$len; $j++) {
                if ($binstr[$j] == "1") {
                    $count += 1;
                }
            }
            if (in_array($count, array(2,3,5,7,11,13,17,19))) {
                $res += 1;
            }
        }
        return $res;
    }
}
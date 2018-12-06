<?php
/*
如果只有$a  $b  $C  三个数组，那么下面这样写也就OK了，但是实现上的需求属性数组是不确定的，
有时是2个数组，有时是3个，有时是4、5、6.....N个，希望写一个无论多少都可以用的方法？

不定数量集合的笛卡尔积
*/

$a=array('黑色', '白色'); 
$b=array('36码', '37码');
$c=array('男款', '女款');

/*
$d = array();
foreach ($a as $_a ){
    foreach ($b as $_b ){
        foreach ($c as $_c ){
            $d[] = $_a.$_b.$_c;
        }
    }
}
print_r($d);

Array
(
    [0] => 黑色36码男款
    [1] => 黑色36码女款
    [2] => 黑色37码男款
    [3] => 黑色37码女款
    [4] => 白色36码男款
    [5] => 白色36码女款
    [6] => 白色37码男款
    [7] => 白色37码女款
)
*/
function Descartes()
{
    $t = func_get_args();

    if (func_num_args() == 1) {
        $t0 = $t[0];
        return call_user_func_array(__FUNCTION__, $t0);
    }

    $a = array_shift($t);
    if (!is_array($a)) $a = array($a);
    $a = array_chunk($a, 1);
    do {
        $r = array();
        $b = array_shift($t);
        if (!is_array($b)) $b = array($b);
        foreach ($a as $p)
            foreach (array_chunk($b, 1) as $q)
                $r[] = array_merge($p, $q);
        $a = $r;
    } while ($t);
    return $r;
}

print_r(Descartes($a, $b, $c));


<?php
 
$arr = array(
    array('a1', 'a2'),
    array('b1', 'b2'),
    array('c1', 'c2')
);
 
fun($arr);
print_r($res);
 
function fun($arr, $tmp = array())
{
    foreach(array_shift($arr) as $v) {
        $tmp[]  = $v;
        if($arr) {
            fun($arr, $tmp);
        }
        else {
            $GLOBALS["res"][]   = implode(" ", $tmp);
        }
        array_pop($tmp);
    }
}

/**
* Descartes
* 笛卡尔积
* @param array(array(1,2,3),array('a','b'))
* @return array('1a','1b','2a','2b','3a','3b')
*/
public function Descartes($array){
    $a = array_pop($array);
    $b = array_pop($array);
    $result = array();
    foreach($a as $av){
        foreach($b as $bv){
            $result[] = $bv.$av;//组合                                                                                                  
        }
    }
    $array[] = $result;//装入
    //判断是否满足递归条件
    if(count($array) > 1){
        $result = Descartes($array);
    }
    return $result;
}
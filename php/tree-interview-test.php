<?php
/**
 * Traverse every node of a tree, print out the path of nodes that contain
 * specific substring, for example, a tree as following
 *
 *           abc
 *         /  |  \
 *      def  gh   ijk  
 *            |
 *           sde
 *
 * for substring 'de', the correct output should be
 *
 *   abc/def
 *   abc/gh/sde
 */
class Node
{
    /**
     * The string of the Node
     *
     * @var string
     */
    protected $str;

    /**
     * Child nodes of this Node
     *
     * @var array of Nodes
     */
    protected $children = array();

    /**
     * constructor
     *
     * @param string $str
     */
    public function  __construct($str) {
        $this->str = $str;
    }

    /**
     * add a child to this node with string
     *
     * @param string $str
     * @return Node
     */
    public function addChild($str) {
        $node = new Node($str);
        $this->addNode($node);
        return $node;
    }

    /**
     * add a child to this node
     * 
     * @param Node $node
     * @return Node 
     */
    public function addNode($node) {
        $this->children[] = $node;
        return $node;
    }

    /**
     * To see if this node contains given substr
     * 
     * @param string $substr
     * @return boolean 
     */
    public function contains($substr) {
        return strstr($this->str, $substr);
    }

    /**
     * Retrieve the string of Node
     * 
     * @return string 
     */
    public function getStr() {
        return $this->str;
    }

    /**
     * Retrieve the children of Node
     * 
     * @return array of Nodes 
     */
    public function getChildren() {
        return $this->children;
    }
}

class Tree
{
    /**
     * Root node of the tree
     *
     * @var Node
     */
    protected $root;

    /**
     * Retrieve the root node of the tree
     *
     * @return Node
     */
    public function getRoot() {
        return $this->root;
    }

    /**
     * Init the tree
     *
     * @return void
     */
    public function init() {
        $this->root = new Node('abc');
        $this->root->addChild('def');
        $this->root->addChild('gh')->addChild('sde');
        $this->root->addChild('ijk');
    }

    /**
     * Traverse every node of the tree
     * 
     * @param Node $node
     * @param string $path
     * @param string $needle
     * @return void
     */
    public function traverse($node, $path, $needle) {
        if (empty($path)) {
            $path = $node->getStr();
        } else {
            $path = $path.'/'.$node->getStr();
        }

        if ($node->contains($needle)) {
            echo $path, PHP_EOL;
        }

        foreach($node->getChildren() as $child) {
            $this->traverse($child, $path, $needle);
        }
    }
}

$tree = new Tree();
$tree->init();
$tree->traverse($tree->getRoot(), '', 'de');

?>
<?php

class Node
{
    /**
     * The string of the Node
     * @var string
     */
    public $str;

    /**
     * Child nodes of this Node
     * @var array of Nodes
     */
    public $children = array();

    /**
     * constructor
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
        return $this; // it's better to return $node
    }

    /**
     * add a child to this node
     * 
     * @param Node $node
     * @return Node 
     */
    public function addNode($node) {
        $this->children[] = $node;
        return $this;
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
}

class Tree
{
    /**
     * Root node of the tree
     * @var Node
     */
    public $root;

    /**
     * Init the tree
     *
     * @return void
     */
    public function init() {
        $this->root = new Node('abc');
        $this->root->addChild('def');
        $node = new Node('gh');
        $this->root->addNode($node);
        $this->root->addChild('ijk');
        $node->addChild('sde');
    }

    /**
     * Traverse every node of the tree
     * 
     * @param Node $node
     * @param string $path
     * @param string $needle
     * @return void
     * @todo explain why static is used with this metohd
     */
    public static function traverse($node, $path, $needle) {
        if (empty($path)) {
            $path = $node->str;
        } else {
            $path = $path.'/'.$node->str;
        }
        if ($node->contains($needle)) {
            echo $path.PHP_EOL;
        }
        foreach($node->children as $child) {
            self::traverse($child, $path, $needle);
        }
    }
}

$tree = new Tree();
$tree->init();

Tree::traverse($tree->root, '', 'de');

?>
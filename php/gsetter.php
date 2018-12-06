<?php

if (count($argv) == 3) {
    require $argv[1];
    generateSettersAndGetters($argv[2]);
}

function generateSettersAndGetters($className)
{
#   $object = new $className();

    $reflect = new ReflectionClass($className);
    $props = $reflect->getProperties();

    foreach ($props as $prop) {

        $propName = $prop->getName();

        /** no setter/getter for field 'id' */
        if (strtolower($propName) == 'id')
            continue;

        /** no setter/getter for inherited fields from parent class */
        if ($prop->getDeclaringClass()->getName() != $className)
            continue;

        generateGetter($propName);
        generateSetter($propName);
    }
}

function generateGetter($propName)
{
    echo "    /**\n";
    echo "     * @return string\n";
    echo "     */\n";
    echo "    public function get", ucfirst($propName), "()\n";
    echo "    {\n";
    echo "        return \$this->$propName;\n";
    echo "    }\n\n";
}

function generateSetter($propName)
{
    echo "    /**\n";
    echo "     * @param \$$propName\n";
    echo "     */\n";
    echo "    public function set", ucfirst($propName), "(\$$propName)\n";
    echo "    {\n";
    echo "        \$this->$propName = \$$propName;\n";
#   echo "        return \$this;\n";
    echo "    }\n\n";
}

<?php
namespace Slim; /** Slim Framework (http://slimframework.com) */

use ArrayIterator;
use Slim\Interfaces\CollectionInterface;

/**
 * Collection
 *
 * This class provides a common interface used by many other
 * classes in a Slim application that manage "collections"
 * of data that must be inspected and/or manipulated
 */
class Collection implements CollectionInterface
{
    protected $data = [];

    /**
     * Create new collection
     * @param array $items Pre-populate collection with this key-value array
     */
    public function __construct(array $items = [])
    {
        foreach ($items as $key => $value) {
            $this->set($key, $value);
        }
    }

    /********************************************************************************
     * Collection interface
     *******************************************************************************/

    /**
     * Set collection item
     *
     * @param string $key   The data key
     * @param mixed  $value The data value
     */
    public function set($key, $value) { $this->data[$key] = $value; }

    /**
     * Get collection item for key
     * @param string $key     The data key
     * @param mixed  $default The default value to return if data key does not exist
     * @return mixed The key's value, or the default value
     */
    public function get($key, $default = null)
    {
        return $this->has($key) ? $this->data[$key] : $default;
    }

    /**
     * Add item to collection
     * @param array $items Key-value array of data to append to this collection
     */
    public function replace(array $items)
    {
        foreach ($items as $key => $value) {
            $this->set($key, $value);
        }
    }

    public function all() { return $this->data; }
    public function keys() { return array_keys($this->data); }

    /**
     * Does this collection have a given key?
     * @param string $key The data key
     * @return bool
     */
    public function has($key) { return array_key_exists($key, $this->data); }
    public function remove($key) { unset($this->data[$key]); }
    public function clear() { $this->data = []; }

    /********************************************************************************
     * ArrayAccess interface
     *******************************************************************************/

    /**
     * Does this collection have a given key?
     * @param  string $key The data key
     * @return bool
     */
    public function offsetExists($key) { return $this->has($key); }

    /**
     * Get collection item for key
     * @param string $key The data key
     * @return mixed The key's value, or the default value
     */
    public function offsetGet($key) { return $this->get($key); }

    /**
     * Set collection item
     * @param string $key   The data key
     * @param mixed  $value The data value
     */
    public function offsetSet($key, $value) { $this->set($key, $value); }

    /**
     * Remove item from collection
     * @param string $key The data key
     */
    public function offsetUnset($key) { $this->remove($key); }

    /**
     * Get number of items in collection
     * @return int
     */
    public function count() { return count($this->data); }

    /********************************************************************************
     * IteratorAggregate interface
     *******************************************************************************/

    /**
     * Get collection iterator
     * @return \ArrayIterator
     */
    public function getIterator() { return new ArrayIterator($this->data); }
}

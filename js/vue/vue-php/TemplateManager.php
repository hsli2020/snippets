<?php // https://github.com/arrilot/vue-templates-in-php

namespace Arrilot\VueTemplates;

class TemplateManager
{
    /**
     * An array of vue components' templates to be injected into html via showTemplates() method
     * @var array
     */
    protected $templates = [];

    /**
     * An array of bundles of vue components' templates to be injected into html via showTemplates() method
     * @var array
     */
    protected $bundles = [];

    /**
     * Path to directory with templates.
     * @var string
     */
    protected $path;

    /**
     * @param string $path
     */
    public function __construct($path)
    {
        $this->path = $path;
    }

    /**
     * Add a template for a component named $name.
     * You can pass some extra data to template using $data array
     *
     * @param string $name
     * @param array $data
     */
    public function addTemplate($name, $data = [])
    {
        if (!isset($this->templates[$name]) || $this->templates[$name] !== false) {
            $this->templates[$name] = compact('name', 'data');
        }
    }

    /**
     * Add several templates at once
     *
     * @param array $names
     * @param array $data
     */
    public function addTemplates($names, $data = [])
    {
        foreach ($names as $name) {
            $this->addTemplate($name, $data);
        }
    }

    /**
     * Add all templates from bundle $name.
     *
     * @param string $name
     * @param array $data
     */
    public function addBundle($name, $data = [])
    {
        if (empty($this->bundles[$name])) {
            return;
        }

        foreach ($this->bundles[$name] as $template) {
            $this->addTemplate($template, $data);
        }
    }

    /**
     * Register new bundle of templates with name $name
     * and array of templates $templates inside it.
     *
     * @param string $name
     * @param array $templates
     */
    public function defineBundle($name, $templates)
    {
        $this->bundles[$name] = $templates;
    }

    /**
     * Print all templates as html.
     */
    public function printTemplates()
    {
        do {
            foreach ($this->templates as $name => $template) {
                if ($template === false) {
                    continue;
                }

                ?><script type="text/x-template" id="vue-<?= htmlspecialchars(str_replace('/', '-', $name), ENT_COMPAT, "UTF-8") ?>-template"><?php
                    $data = $template['data'];
                    require $this->path . '/' . $name . '.php';
                ?></script><?php

                // mark template as already printed.
                $this->templates[$name] = false;
            }

            // if we add even more templates in templates, they must be printed as well.
            $newTemplates = array_filter($this->templates, function($template) {
                return $template !== false;
            });
        } while (count($newTemplates) > 0);
    }
}

<?php

use Phalcon\Acl;
use Phalcon\Acl\Role;
use Phalcon\Acl\Resource;
use Phalcon\Acl\Adapter\Memory as AclEngine;

//example1();
example2();

function example1()
{
    $acl = new AclEngine();

    $acl->setDefaultAction(Acl::DENY);

    $acl->addRole(new Role('Users'));
    $acl->addRole(new Role('Managers'), 'Users');

    $acl->addResource(new Resource('Articles'), ['search', 'update']);

    $acl->allow('Users', 'Articles', 'search'); // "Managers" inherit from "Users" its accesses

    var_dump($acl->isAllowed('Users',    'Articles', 'search'));
    var_dump($acl->isAllowed('Managers', 'Articles', 'search'));

    var_dump($acl->isAllowed('Users',    'Articles', 'update'));
    var_dump($acl->isAllowed('Managers', 'Articles', 'update'));

    //print_r($acl);
}

function example2()
{
    $acl = new AclEngine();

    $acl->setDefaultAction(Acl::DENY);

    $acl->addRole('user');
    $acl->addRole('admin', 'user');
    $acl->addRole('developer', 'admin');

    $acl->addResource('tickets', ['list', 'open', 'close']);

    $acl->allow('user', 'tickets', 'open');

    var_dump($acl->isAllowed('user', 'tickets', 'open'));
    var_dump($acl->isAllowed('admin', 'tickets', 'open'));
    var_dump($acl->isAllowed('developer', 'tickets', 'open'));
}

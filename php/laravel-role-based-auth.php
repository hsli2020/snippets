<?php
# Table
/*
users
	- id
	- name
	- email
	- password

roles
	- id
	- name
	- label
	
permissions // abilities
	- id
	- name
	- label

role_permission
	- role_id
	- permission_id
	UNIQUE role_id + permission_id

user_role
	- user_id
	- role_id
	UNIQUE user_id + role_id
*/

# Model

class User extends Model
{
	public function roles()
	{
		return $this->belongsToMany(Role::class);
	}

    public function assignRole($role)
    {
        if (is_string($role)) {
            $role = Role::whereName($role)->firstOrFile();
        }

        $this->roles()->save($role);
       #$this->roles()->sync($role, false); // create or update
    }

    public function permissions()
    {
        return $this->roles->map->permissions->flatten()->pluck('name')->unique();
    }
}

class Role extends Model
{
	public function permissions()
	{
		return $this->belongsToMany(Permission::class);
	}

    public function allowTo($perm)
    {
        if (is_string($perm)) {
            $perm = Permission::whereName($perm)->firstOrFile();
        }

        $this->permissions()->save($perm);
       #$this->permissions()->sync($perm, false); // create or update
    }
}

class Permission extends Model
{
	public function roles()
	{
		return $this->belongsToMany(Role::class);
	}
}

# Tinker

$user = User::find(1);

$role = Role::firestOrCreate([ 'name' => 'moderator' ]);

$perm = Permission::firstOrCreate([ 'name' => 'edit-form' ]);

$role->allowTo($perm);

$user->assignRole($role);

// $user->roles
// $user->roles[0]->permissions
// $user->permissions()  // $user->roles->map->permissions->flatten()->pluck('name')->unique();

class AuthServiceProvide extends ServiceProvider
{
    public function root()
    {
        $this->registerPolicies();

        Gate::before(function ($user, $perm) {
            return $user->permissions()->contains($perm);
        });
    }
}

# index.blade.php

@can('edit-form')
    <p>You have perm to see this</p>
@endcan

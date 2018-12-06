<?php namespace App\Models;

use Illuminate\Auth\Authenticatable;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Auth\Passwords\CanResetPassword;
use Illuminate\Contracts\Auth\Authenticatable as AuthenticatableContract;
use Illuminate\Contracts\Auth\CanResetPassword as CanResetPasswordContract;

class User extends Model implements AuthenticatableContract, CanResetPasswordContract 
{
	use Authenticatable, CanResetPassword;

	protected $table = 'users';

	/** The attributes excluded from the model's JSON form. */
	protected $hidden = ['password', 'remember_token'];

	public function role()      { return $this->belongsTo('App\Models\Role'); }
	public function posts()     { return $this->hasMany('App\Models\Post'); }
	public function comments()  { return $this->hasMany('App\Models\Comment'); }

	public function isAdmin()   { return $this->role->slug == 'admin'; }
	public function isNotUser() { return $this->role->slug != 'user'; }

	public function accessMediasAll()    { return $this->isAdmin(); }
	public function accessMediasFolder() { return $this->isNotUser(); }
}

class Role extends Model
{
	public function users() { return $this->hasMany('App\Models\User'); }
}

use App\Presenters\DatePresenter;

class Post extends Model
{
	use DatePresenter;

	public function user()     { return $this->belongsTo('App\Models\User'); }
	public function tags()     { return $this->belongsToMany('App\Models\Tag'); } 
	public function comments() { return $this->hasMany('App\Models\Comment'); }
}

class Comment extends Model
{
	use DatePresenter;

	public function user() { return $this->belongsTo('App\Models\User'); }
	public function post() { return $this->belongsTo('App\Models\Post'); }
}

class Tag extends Model
{
	public function posts() { return $this->belongsToMany('App\Models\Post'); }
}

class PostTag extends Model
{
	protected $table = 'post_tag';
	public $timestamps = false;
}

<?php

/**
 * Idiomatic way to define a class in PHP 8.x
 */
final class User
{
    /**
     * define all properties in constructor
     */
    public function __construct(
        public readonly string $username,
        public readonly string $email,
        #[\SensitiveParameter]
        public readonly string $password
    )
    {
    }

    /**
     * use a static method as factory
     */
    public static function from($data)
    {
        return new self(
            $data['username'],
            $data['email'],
            $data['password']
        );
    }

    public function toArray()
    {
        return [
            'username' => $this->username,
            'email'    => $this->email,
            'password' => $this->password,
        ];
    }
}

$data = [
    'username' => 'John Doe',
    'email'    => 'johndoe@email.com',
    'password' => '123456',
];

$user = User::from($data);
print_r($user->toArray());

<?php

require_once __DIR__ . '/vendor/autoload.php';

use \Symfony\Component\HttpFoundation\Response;
use \Symfony\Component\HttpClient\HttpClient;

$client = HttpClient::create();

if (rand(0, 100) === 0) {
    error_log("HASH");
    password_hash('password', PASSWORD_BCRYPT, ['cost' => 12]);
} else {
    $client->request('GET', 'some-url-you-can-put-load-on');
}

$response = new Response('OK');
$response->send();

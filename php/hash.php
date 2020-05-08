<?php

require_once __DIR__ . '/vendor/autoload.php';

use \Symfony\Component\HttpFoundation\JsonResponse;
use \Symfony\Component\HttpFoundation\Response;
use \Symfony\Component\HttpClient\HttpClient;

$client = HttpClient::create();

$hash = password_hash('password', PASSWORD_BCRYPT, ['cost' => 12]);

$response = new JsonResponse(['hash' => $hash]);
$response->send();

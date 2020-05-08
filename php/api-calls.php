<?php

require_once __DIR__ . '/vendor/autoload.php';

use \Symfony\Component\HttpFoundation\Response;
use \Symfony\Component\HttpClient\HttpClient;

$client = HttpClient::create();

$client->request('GET', 'some-url-you-can-put-load-on');

$response = new Response('OK');
$response->send();

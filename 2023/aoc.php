<?php
require "vendor/autoload.php";

if($argc <= 1) {
    echo "Please choose a day to run";
    die;
}

$day = (int) $argv[1];

if(!$day){
    echo "Please provide a number for the day";
    die;
}

include "./day-$day/day-" . $day . ".php";

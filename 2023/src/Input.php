<?php

namespace Aoc\TwentyThree;

class Input {
    public string $path;
    public string $line;
    private $handle;

    public function __construct(){
    }

    public function read($path){
        $this->path = $path;
        $handle = fopen($this->path, "r");
        if(!$handle) {
            throw new \RuntimeException("Cant read $path");
        }
        $this->handle = $handle;
    }

    public function next(){
        $this->line = fgets($this->handle);
        if($this->line === false){
            return null;
        }

        return $this->line;
    }

}

<?php
use Aoc\TwentyThree\Input;
use Illuminate\Support\Collection;

dayOne();

/**
 * @return void 
 * @throws Exception 
 */
function dayOne() {
    $input = new Input;
    $input->read("./day-1/day-1.input.txt");
    $col = new Collection;
    while ($line = $input->next()) {
        $col->add(
            implode(
                '',
                firstLastNumeric($line)
            )
        );
    }
    print_r($col->sum());
}

/**
 * @param mixed $line 
 * @return string[] 
 * @throws Exception 
 */
function firstLastNumeric($line) {
    return [
        findFirst($line),
        findLast($line)
    ];
}

/**
 * @param string $line 
 * @return string 
 * @throws Exception 
 */
function findFirst(string $line) {
    $pointer = 0;
    while ($pointer < strlen($line)) {
        $char = substr($line, $pointer, 1);
        if (is_numeric($char)) {
            return $char;
        }
        if ($number = alpha_number($line, $pointer)) {
            return strval($number);
        }

        $pointer++;
    }
    throw new \Exception("Couldnt find the first Integer on line: $line");
}

/**
 * @param string $line 
 * @return string 
 * @throws Exception 
 */
function findLast($line): string {
    $pointer = strlen($line) - 1;
    while ($pointer >= 0) {
        $char = substr($line, $pointer, 1);
        if (is_numeric($char)) {
            return $char;
        }
        if ($number = alpha_number($line, $pointer)) {
            return strval($number);
        }

        $pointer--;
    }
    throw new \Exception("Couldnt find the last Integer on line: $line");
}

/**
 * @param string $line 
 * @param int $pointer 
 * @return int|null 
 */
function alpha_number(string $line, int $pointer): int|null {
    $numbers = new Collection([ "one" => 1, "two" => 2, "three" => 3, "four" => 4, "five" => 5, "six" => 6, "seven" => 7, "eight" => 8, "nine" => 9, ]);
    $alpha_number = $numbers->first(function($val, $key) use ($line, $pointer){
       return $key == substr($line, $pointer, strlen($key));
    });

    return $alpha_number;
}

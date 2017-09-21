(ns wedge2.core-test
  (:require [clojure.test :refer :all]
            [wedge2.core :refer :all]
            [clj-time.core :as t]))

(deftest test-sort-movements
	(let [
		input  [{:date (t/local-date 2017 1 2)} {:date (t/local-date 2017 3 2)} {:date (t/local-date 2014 9 2)}]
		output [{:date (t/local-date 2014 9 2)} {:date (t/local-date 2017 1 2)} {:date (t/local-date 2017 3 2)}]]
    (is (= output (sort-movements input)))))

(deftest test-balances-1
	(let [
		input  []
		output {}]
    (is (= output (balances input)))))

(deftest test-balances-2
	(let [
		input  [{:account "A" :date (t/local-date 2017 1 2) :value 3/4} {:account "B" :date (t/local-date 2017 3 2) :value 20} {:account "B" :date (t/local-date 2014 9 2) :value -1/4}]
		output {"A" [{:date (t/local-date 2017 1 2) :value 3/4}] "B" [{:date (t/local-date 2014 9 2) :value -1/4} {:date (t/local-date 2017 3 2) :value 79/4}]}]
    	(is (= output (balances input)))))
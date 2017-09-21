(ns wedge2.core-test
  (:require [clojure.test :refer :all]
            [wedge2.core :refer :all]
            [clj-time.core :as t]))

(deftest test-sort-movements-1
	(let [
		input  []
		output []]
    (is (= output (sort-movements input)))))

(deftest test-sort-movements-2
	(let [
		input  [{:date (t/local-date 2017 1 2)}]
		output [{:date (t/local-date 2017 1 2)}]]
    (is (= output (sort-movements input)))))

(deftest test-sort-movements-3
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
		input  [
			{:account "A" :date (t/local-date 2017 1 2) :value 3/4}]
		output {
			["A" nil] [{:date (t/local-date 2017 1 2) :value 3/4}]}]
    (is (= output (balances input)))))

(deftest test-balances-3
	(let [
		input  [
			{:account "A" :date (t/local-date 2017 1 2) :value 3/4}
			{:account "B" :date (t/local-date 2017 3 2) :value 20}
			{:account "B" :date (t/local-date 2014 9 2) :value -1/4}]
		output {
			["A" nil] [{:date (t/local-date 2017 1 2) :value 3/4}]
			["B" nil] [{:date (t/local-date 2014 9 2) :value -1/4} {:date (t/local-date 2017 3 2) :value 79/4}]}]
    	(is (= output (balances input)))))

(deftest test-balances-4
	(let [
		input  [
			{:account "A" :asset "BRL" :date (t/local-date 2017 1 2) :value 3/4}
			{:account "A" :asset "BRL" :date (t/local-date 2017 1 2) :value -1/4}
			{:account "A" :asset "USD" :date (t/local-date 2017 1 2) :value 314/100}
			{:account "B" :asset "BRL" :date (t/local-date 2017 3 2) :value 20}
			{:account "B" :asset "BRL" :date (t/local-date 2014 9 2) :value -1/4}]
		output {
			["A" "BRL"] [{:date (t/local-date 2017 1 2) :value 2/4}]
			["A" "USD"] [{:date (t/local-date 2017 1 2) :value 314/100}]
			["B" "BRL"] [{:date (t/local-date 2014 9 2) :value -1/4} {:date (t/local-date 2017 3 2) :value 79/4}]}]
    	(is (= output (balances input)))))
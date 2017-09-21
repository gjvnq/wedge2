(ns wedge2.core
	(:gen-class)
	(:require
		[cheshire.core :as chess]
		[clj-time.core :as t]))

(load "domain")

(defn -main
	"I don't do a whole lot ... yet."
	[& args]
	(println "Hello, World!" args))

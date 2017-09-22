(defproject wedge2 "0.1.0-SNAPSHOT"
  :description "FIXME: write description"
  :url "http://example.com/FIXME"
  :license {:name "LGPL"
            :url "https://www.gnu.org/licenses/lgpl-3.0.en.html"}
  :dependencies [[org.clojure/clojure "1.8.0"] [cheshire "5.8.0"] [clj-time "0.14.0"]]
  :main ^:skip-aot wedge2.core
  :target-path "target/%s"
  :profiles {:uberjar {:aot :all}})

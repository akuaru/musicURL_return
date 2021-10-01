個人的に作成したものです。フォルダ名が適当なので読み替えるか書き換えることを推奨します。<br>
musicURL_returnをカレントディレクトリにしてdocker-compose up -d を行うとMYSQLのDBが立ち上がります。<br>
その後、go run main.go を行いhttp://localhost:8080/users にアクセスするとDB内のデータを返します。<br>
参考文献：https://g-hyoga.hatenablog.com/entry/2018/01/29/003033<br>
https://medium.com/@taka.abc.hiko/gorm%E3%81%A8%E3%81%84%E3%81%86orm%E3%81%8C%E8%B6%85%E7%B5%B6%E4%BE%BF%E5%88%A9%E3%81%A0%E3%81%A3%E3%81%9F%E4%BB%B6-8d279489c38f
解説かいておく

fmt
Printfは、Fprintfを呼び出している引数にos.Stdoutを渡している

Fprintfの引数をみると、os.Stdoutはos.Writerとして扱われる

io.Writerはinterface

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

どこぞにデータを置くinterfaceとして、一般的に使われるもの

これを使うと、greetをどこぞに出力するものとして使う

すなわち、greetがwriterを受け取るようにし、testのwriterにはbytes.Bufferを使う
bytes.Bufferはio.Writer interfaceを実装しているため

また、greet内ではstdoutに書き出すfmt.Printfではなく、
writerに書き出すfmt.Fprintfを用いる

greetはwriterで指定された場所に値を書き出すなにかとなる


>func Greet(writer *bytes.Buffer, name string) {

としておくと、writerとしてbytes.Bufferしか受け取れなくなるため

interfaceとしてio.Writerを実装しているもの、(bytes.Buffer, os.Stdoutなど)を
受け取れるように引数を変える


テスト可能にすることを目的として
injecting a dependencyを行うことで、値の書き込み先を制御することできた

テストする際に、dependenciesが強いと（DBから書き込んで、サービスに返す処理がまとまっているなど）、テストを実装するのは難しくなる。DIはDBの依存部分をinjectするなどし、
mock outすることを可能にする

関心の分離、データがどこに返されるか、データがどのように生成、取得されているかどうか、を分離する
メソッドの責務の範囲をを細分化する

別なコンテキストで、コードを再利用できる。別なdependencyをinjectionできるので
package main
import (
  "fmt"
  "os"
  "io/ioutil"
  "sort"
  "flag"
  "github.com/Songmu/prompter"
)

type FileInfos []os.FileInfo

//sort を使用するために以下の3つのインタフェースを満たす必要がある
func (fi FileInfos) Len() int {
  return len(fi)
}
func (fi FileInfos) Swap(i, j int) {
  fi[i], fi[j] = fi[j], fi[i]
}
func (fi FileInfos) Less(i, j int) bool {
  return fi[i].ModTime().Unix() < fi[j].ModTime().Unix()
}

func main(){
  // 引数の数の確認
  // if len(os.Args) != 2 {
  //   fmt.Println("argument error.")
  //   os.Exit(1)
  // }

  // 引数を格納
  //dir_path := os.Args[1]

  // 引数の取得にflagを利用
  var arg string
  flag.StringVar(&arg, "f", "", "directory path")
  flag.Parse()

  // カレントディレクトリのパスを取得
  var curDir, _ = os.Getwd()
//  curDir += "/"

  // -fオプションに何も入力されていない場合、カレントディレクトリ
  if arg == "" {
    arg = curDir
  }

  // -fオプションされずにパスが入力されていた場合
  if flag.Arg(0) != "" {
    arg = flag.Arg(0)
  }

  // 処理対象のパスをセット
  dir_path := arg
  if !(prompter.YN("Diretcory path OK? " + dir_path, false)) {
      fmt.Println("stop process.")
      os.Exit(0)
  }

  // パスの情報を取得
  fInfo, err := os.Stat(dir_path)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  // ディレクトリでなければエラー
  if !(fInfo.IsDir()) {
    fmt.Println("Error: Please input directory path.")
    os.Exit(1)
  }

  // ディレクトリ配下のファイル一覧取得
  fileInfos,err := ioutil.ReadDir(dir_path)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // ファイル一覧表示(sort前)
  for index, fileInfo := range fileInfos {
    fmt.Printf("index:%d name:%s\t", index, fileInfo.Name())
    fmt.Printf("mtime:%s\n", fileInfo.ModTime())
  }

  //sort処理
  sort.Sort(FileInfos(fileInfos))

  fmt.Println("\nafter sort.")
  // ファイル一覧表示(sort後)
  fmt.Println(fileInfos[0].Name())
  for index, fileInfo := range fileInfos {
    fmt.Printf("index:%d name:%s\t", index, fileInfo.Name())
    fmt.Printf("mtime:%s\n", fileInfo.ModTime())
  }

}

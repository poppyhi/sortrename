package main
import (
  "fmt"
  "os"
  "io/ioutil"
  "sort"
  "flag"
  "github.com/Songmu/prompter"
  "path/filepath"
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
  // 引数の取得にflagを利用
  var arg string
  flag.StringVar(&arg, "f", "", "directory path")
  flag.Parse()

  // カレントディレクトリのパスを取得
  var curDir, _ = os.Getwd()

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
  // パスの最後が"/"でなかった場合、追加
  if dir_path[len(dir_path)-1:] != "/" {
    dir_path += "/"
  }
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

  //sort処理
  sort.Sort(FileInfos(fileInfos))

  // ファイル一覧表示(sort後)
  for index, fileInfo := range fileInfos {
    // 001からとしたいので、+1
    var index_str string = fmt.Sprint(index+1)

    // 000 の3桁にする
    if index < 9 {
      index_str = "00" + index_str
    } else if index < 99 {
      index_str = "0" + index_str
    }

    // 拡張子を取得し、リネームするファイル名につける
    pos := filepath.Ext(fileInfo.Name())
    index_str += pos

    // リネーム
    if err := os.Rename(dir_path+fileInfo.Name(), dir_path+index_str ); err != nil {
      fmt.Println(err)
    }
  }

}

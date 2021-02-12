# cloud-function-avoid-cli
Avoid Cloud Functions and continue to use Firebase projects for free.  (fantasy...)


## Github Actions
> [How to use Github Actions?](https://developer.yukimonkey.com/article/20200422/)

### Local Github Actions
> Install
> ```shell
> brew install nektos/tap/act
> ```
> 
> Run
> ```shell
> act push --secret-file .secrets
> ```
> 
> 

## Golang Cobra Cli 
> [Cobraの使い方とテスト](https://text.baldanders.info/golang/using-and-testing-cobra/)
> 
> [Cobra でテストしやすい CLI を構成する](https://zenn.dev/spiegel/articles/20201018-cli-with-cobra-and-golang)

## How to use Firebase Admin SDK with Golang?
> [GolangでFirebase Admin SDKの認証周りを試しました](https://qiita.com/nisitanisubaru/items/3ff4e0b08b20700f408c)

## Setting up the test environment
> absolute path
> ```shell
> find `pwd` -name "firebase_adminsdk.json"
> ```
> 
> export google-application-credentials
> ```shell
> export GOOGLE_APPLICATION_CREDENTIALS=${absolute path}
> ```
# gflac
得到周杰伦的flac的音乐

[安装使用视频讲解 youtube](https://www.youtube.com/watch?v=RYdAtgvcpzY)

[安装使用视频讲解 bilibili](https://www.bilibili.com/video/BV1qk4y1z77K/)


当你下载了文件后主义要给权限 `chmod +x ./gflac`
```bash
 gflac -num 10 -peo 周杰伦 -coo flac2020
```
> 因为此网站有可能会限制登陆 如果你遇到无法下载的情况你可以打开这个网站 https://www.52flac.com/  随便找到一个音乐的下载的地方
这个地方如果可以下载会出现百度网盘的地址，如果不能就说明你需要扫码了（上面自动会出现扫码的二维码）这个时候你只需要扫码关注 然后回复
验证码就可以得到 验证码然后你在 -coo 后面加上这个验证码即可，本人跟这个网站没有任何的利益交集，你可以扫码后立马取消关注，反正无所谓。

**经过测试 貌似 flac2020 有可能在2020年能用一年，也就是说这个网站技术一般，并没有做动态的数据验证机制。**

同样的如果你想得到比如王力宏的音乐，只需要将`-pop 周杰伦`改成 `-pop 王力宏`

理论上你不实用姓名搜索，搜索具体的歌名也可以出现。
## 下载

```bash
git clone https://github.com/shgopher/gflac

cd gflac

go build

```


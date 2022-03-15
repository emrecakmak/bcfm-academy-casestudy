
# bcfm-academy-casestudy

Bu repo Bestcloudfor.me case study projesini içerir.



## API Kullanımı

#### Name ve surname bilgisini döner.

```https
  GET localhost:8080
```

#### Şehrin sıcaklık derecesini döner

```https
  GET localhost:8080/temperature=?city=${city}
  
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `city`      | `string` | İstenilen şehir girilmeli |

  

## Bilgisayarınızda Çalıştırın

Projeyi klonlayın

```bash
  git clone https://github.com/emrecakmak/bcfm-academy-casestudy
```

Proje dizinine gidin

```bash
  cd bcfm-academy-casestudy
```

docker kullanarak build alın

```docker
docker build -t bcfm-case:latest . 
```
Containerı çalıştırın
```docker
docker run -d -p 8080:8080 bcfm-case
```

  
## Ekran Görüntüleri

![name surname](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Ffcb65971-9fe3-4058-9253-ee27e6895688%2FUntitled.png?table=block&id=209a07dd-2b72-4ebf-912f-4d8b1f2c229c&spaceId=37eed0df-5489-4147-b89a-b5b1076cbd31&width=2000&userId=131eab5a-b4fa-40bf-90e1-f6977af8faad&cache=v2)


  
![temperature](  https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Fd7acc4c0-5b08-4638-bf58-01c78fe8e3bc%2FUntitled.png?table=block&id=fc4a664c-ee78-4f16-aaa7-d2413720dee0&spaceId=37eed0df-5489-4147-b89a-b5b1076cbd31&width=2000&userId=131eab5a-b4fa-40bf-90e1-f6977af8faad&cache=v2)

## AWS Demo Link

http://ec2-3-125-40-70.eu-central-1.compute.amazonaws.com:8080
http://ec2-3-125-40-70.eu-central-1.compute.amazonaws.com:8080/temperature?city=Ankara

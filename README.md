# 젠킨스 Publish Over SSH 설정 바꾸기

* 영업용으로 들고다니는 테스트서버는 아이피가 수시로 바뀐다.
* 이 서버에 작업물을 수시로 올리기도 해야 해서 젠킨스 플러그인에 물린 아이피 설정도 매번 맞춰줘야 한다.
* 젠킨스의 SSH 플러그인은 `XML 1.1` 포맷이다.
* `Go`에서 `XML 1.0`만 지원하고 `XML 1.1`은 지원하지 않는다.
* 그래서 그냥 XML을 텍스트파일로 열고 정규표현식으로 치환하는 걸로 때우기로 함

끝.
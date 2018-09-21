// 代码中会兼容本地 service mock 以及部署站点的静态数据
export default {
  'POST /api/signin': (req, res) => {
    const { password, username } = req.body;
    if (password === '12345' && username === 'user001') {
      res.send({
          code: 200,
          expire: "2020-09-21T17:08:50Z",
          token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzc1NDk3MzAsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTUzNzU0NjEzMH0.arYyHXJF69U3a6RcLxxQvyUcf7YetV2B86XyCGnkWrs",
      });
      return;
    }
    res.send({
      code: 401,
    });
  },
};

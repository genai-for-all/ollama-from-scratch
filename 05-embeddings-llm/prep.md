I will store several vectors (embeddings) like the following in a key-value database (like redis)

{"embedding":[0.2349054217338562,0.07909717410802841,0.21808242797851562,-0.3499607741832733,-0.38203203678131104,-0.4509012699127197,1.0033316612243652,-0.019546981900930405,0.38844552636146545,0.4179795980453491,-0.15826360881328583,0.021358396857976913,0.03496265411376953,0.28549647331237793,-0.16837343573570251,-0.6167565584182739,0.12701740860939026,-0.511426568031311,0.45987164974212646,-0.18750551342964172]}

The store format will be:

{key: <a unique key>, "embedding":[<array of value>]}

At the end of the day I will have a liste like this one:

[
    {key: "AZSX", "embedding":[1.0, 4.0, 0.2]},
    {key: "DCXZ", "embedding":[1.2, 4.3, 0.2]},
    {key: "WXCV", "embedding":[0.0, 1.0, 3.2]},
    {key: "DFKI", "embedding":[-0.5, 1.0, 3.2]},
    {key: "ZERT", "embedding":[0.0, 1.3, 3.2]},
]

If I have a new embedding (vector) like {key: "WWWQ", "embedding":[-0.5, 1.3, 5.2]} how can I calculate the cosine distance between this embedding and each of the other embeddings

# server tools & grpc
from concurrent.futures import ThreadPoolExecutor
from wine_varieties_pb2 import WineReviewResponse
from wine_varieties_pb2_grpc import WineClassifierServiceServicer, add_WineClassifierServiceServicer_to_server
import logging
import grpc
import sklearn

# model construction and selection
from sklearn.svm import LinearSVC
from sklearn.model_selection import train_test_split

# feature processing
from sklearn.feature_extraction.text import CountVectorizer
from sklearn.feature_extraction.text import TfidfTransformer

# data utilities
from collections import Counter
import numpy as np
import pandas as pd

class WineClassifierServer(WineClassifierServiceServicer):
    def __init__(self):
        wine_df = pd.read_csv('data/winemag-data-130k-v2.csv')

        # map top 10 wine varieties to ints in wine_df
        counter = Counter(wine_df['variety'].tolist())
        self.top_10_varieties = {i[0]: idx for idx, i in enumerate(counter.most_common(10))}
        logging.info(self.top_10_varieties)
        wine_df = wine_df[wine_df['variety'].map(lambda x: x in self.top_10_varieties)]

        # convert description and variety to lists
        description_list = wine_df['description'].tolist()
        varietal_list = [self.top_10_varieties[i] for i in wine_df['variety'].tolist()]
        varietal_list = np.array(varietal_list)

        # learn vocabulary and return document-term matrix
        self.count_vect = CountVectorizer()
        x_train_counts = self.count_vect.fit_transform(description_list)

        tfidf_transformer = TfidfTransformer()
        x_train_tfidf = tfidf_transformer.fit_transform(x_train_counts)

        train_x, test_x, train_y, test_y = train_test_split(x_train_tfidf, varietal_list, test_size=0.2)

        # build and fit linear SVC, then make predictions
        self.lsvc = LinearSVC().fit(train_x, train_y)
        y_score = self.lsvc.predict(test_x)

        n_right = 0
        for i in range(len(y_score)):
            if y_score[i] == test_y[i]:
                n_right += 1

        print("Accuracy: %.2f%%" % ((n_right/float(len(test_y)) * 100)))

    def GetWineVariety(self, request, context):
     logging.info(request.review)
    # transform review to document-term matrix
     x_train_counts = self.count_vect.transform([request.review])

     tfidf_transformer = TfidfTransformer()
     x_tfidf = tfidf_transformer.fit_transform(x_train_counts)

     prediction = self.lsvc.predict(x_tfidf)
     for variety, idx in self.top_10_varieties.items():
        if prediction[0] == idx:
            print(variety)
            resp = WineReviewResponse(variety=variety)
     return resp
    

#starting the gRPC  server

if __name__ == '__main__':
    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
    )
    server = grpc.server(ThreadPoolExecutor())
    add_WineClassifierServiceServicer_to_server(WineClassifierServer(), server)
    port = 8080
    # TODO
    # those setups should not be hardcoded,rather it should in a configuration file like  .json
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    logging.info('server ready on port %r', port)
    server.wait_for_termination()
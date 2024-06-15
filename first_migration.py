
import pandas as pd
from sqlalchemy import create_engine

data = pd.read_csv("C:/Users/joaco/Documents/store_zara.csv")

array = data["images"]

array = array.apply(lambda x: x.replace("[", ""))

array = array.apply(lambda x: x.replace("]", ""))

array = array.apply(lambda x: x.replace(' ', ""))
array = array.apply(lambda x: x.replace("'", ""))
array = array.apply(lambda x: x.split(","))

engine = create_engine('postgresql://root:secret@localhost:5432/root')

data["images"] = array

data.to_sql('zara', engine)

data.to_csv("C:/Users/joaco/Documents/zara.csv" )

print("Done!")

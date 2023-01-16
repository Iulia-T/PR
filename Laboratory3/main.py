from flask import Flask
from flask_restful import Api, Resource, abort, reqparse

app = Flask(__name__)
api = Api(app)

food_put_args = reqparse. RequestParser()
food_put_args.add_argument("name", type=str, help="Name of the food is required", required=True)
food_put_args.add_argument("preparation-time", type=int, help="Preparation time of the food is required", required=True)
food_put_args.add_argument("complexity", type=int, help="Complexity of the food is required", required=True)
food_put_args.add_argument("cooking-apparatus", type=str, help="Cooking apparatus neede to prepare the food is required", required=True)

foods = {}

def abort_if_food_id_doesnt_exist(food_id):
    if food_id not in foods:
        abort(404, message = "Food id isn't valid...")
        

class food(Resource):
    def get(self, food_id):
        abort_if_food_id_doesnt_exist(food_id)
        return foods[food_id]
    
    def put(self, food_id):
        args = food_put_args.parse_args()
        foods[food_id] = args
        return foods[food_id], 201 #created
    
    def delete(sef, food_id):
        abort_if_food_id_doesnt_exist(food_id)
        del foods[food_id]
        return '', 204 
    
    #def patch(sef, food_id):
        #abort_if_food_id_doesnt_exist(food_id)
        

api.add_resource(food, "/food/<int:food_id>")

if __name__ == "__main__":
    app.run(debug=True)
    

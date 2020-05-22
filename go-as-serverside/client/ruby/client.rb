$LOAD_PATH.push File.join(File.dirname(__FILE__), '../../pb_ruby')

require 'grpc'
require 'user_services_pb'

def main
  stub = Pb::User::Stub.new('localhost:9999', :this_channel_is_insecure)
  response = stub.get_user_info(Pb::UserRequest.new(name: 'Guobin'))
  puts "Received user name: #{response.name}"
end

main

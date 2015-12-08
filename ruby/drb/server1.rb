require "drb"

URI = "druby://localhost:55331"
class Server
  attr_reader :history

  def initialize
    @history = [].extend(DRb::DRbUndumped)
  end

  def multiply(x, y)
    (x * y).tap do |res|
      puts "#{x} * #{y} = #{res}"
      history.push res
    end
  end

  def perform(*args, &block)
    puts "Performing with arguments: #{args}"
    history[history.size] = block.call(args)
  end
end

DRb.start_service(URI, Server.new)
DRb.thread.join

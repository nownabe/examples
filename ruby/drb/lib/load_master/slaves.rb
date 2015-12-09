require "drb"
require "parallel"

module LoadMaster
  class Slaves
    def initialize(drb_uris)
      @slaves = drb_uris.map { |uri| DRbObject.new_with_uri(uri) }
    end

    def count
      @slaves.count
    end

    def load(test)
      Parallel.map(@slaves, in_threads: count) do |slave|
        slave.generate_loader(test).execute
      end
    end
  end
end

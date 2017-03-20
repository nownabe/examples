import com.google.auth.oauth2.ServiceAccountCredentials;
import com.google.cloud.bigquery.*;

import java.io.FileInputStream;
import java.io.IOException;

public class Main {
    private static final String DATASET = "yourdataset";
    private static final String TABLE = "yourtable";

    public static void main(String args[]) throws IOException {
        BigQuery bigQuery;
        if (args.length > 0) {
            bigQuery = getClientWithJsonKey(args[0]);
        } else {
            bigQuery = BigQueryOptions.getDefaultInstance().getService();
        }

        Table table = bigQuery.getTable(DATASET, TABLE);
        if (table != null) {
            System.out.println("Found!");
        } else {
            System.err.println("Not found");
        }
    }

    private static BigQuery getClientWithJsonKey(String key) throws IOException {
        return BigQueryOptions.newBuilder()
                .setCredentials(ServiceAccountCredentials.fromStream(new FileInputStream(key)))
                .build()
                .getService();
    }
}


import com.google.auth.oauth2.ServiceAccountCredentials;
import com.google.cloud.Page;
import com.google.cloud.storage.Blob;
import com.google.cloud.storage.Bucket;
import com.google.cloud.storage.Storage;
import com.google.cloud.storage.StorageOptions;

import java.io.*;
import java.util.Iterator;

public class GCSList {
    private static final String BUCKET = "mybucket";
    private static final String PREFIX = "dir1/dir2/";

    public static void main(String args[]) throws IOException {
        Storage storage;
        if (args.length > 0) {
            storage = getStorageFromJsonKey(args[0]);
        } else {
            storage = StorageOptions.getDefaultInstance().getService();
        }

        Bucket bucket = storage.get(BUCKET);
        Storage.BlobListOption option = Storage.BlobListOption.prefix(PREFIX);
        Page<Blob> blobs = bucket.list(option);
        Iterator<Blob> blobIterator = blobs.iterateAll();

        while (blobIterator.hasNext()) {
            System.out.println(blobIterator.next().getName());
        }
    }

    private static Storage getStorageFromJsonKey(String key) throws IOException {
        return StorageOptions.newBuilder()
                .setCredentials(ServiceAccountCredentials.fromStream(new FileInputStream(key)))
                .build()
                .getService();
    }
}

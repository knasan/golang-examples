# disk-info2

I used a library in disk-info, I only needed hard drive information, so I looked for an alternative.

With the Linux command

```sh
    lsblk -I 8,179,252,259 -f -J -o fstype, name, label, size, mountpoint
```

```console
{
   "blockdevices": [
      {"fstype":null, "name":"nvme0n1", "label":null, "size":"953,9G", "mountpoint":null,
         "children": [
            {"fstype":"vfat", "name":"nvme0n1p1", "label":null, "size":"512M", "mountpoint":"/boot/efi"},
            {"fstype":"ext4", "name":"nvme0n1p2", "label":null, "size":"23,3G", "mountpoint":"/"},
            {"fstype":"ext4", "name":"nvme0n1p3", "label":null, "size":"9,3G", "mountpoint":"/var"},
            {"fstype":"swap", "name":"nvme0n1p4", "label":null, "size":"977M", "mountpoint":"[SWAP]"},
            {"fstype":"ext4", "name":"nvme0n1p5", "label":null, "size":"1,9G", "mountpoint":"/tmp"},
            {"fstype":"ext4", "name":"nvme0n1p6", "label":null, "size":"918G", "mountpoint":"/home"}
         ]
      }
   ]
}
```

you can have the hard disk information output as JSON, which is perfect for processing with Go.

This should not be an example, it served me more as a playground and test until I used it in one
of my programs in a slightly better spelling and tidy code.
I think the method is good, so I definitely want to publish it in my example list as an inspiration.
